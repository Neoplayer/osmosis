use crate::msg::{Channel, QuotaMsg};
use crate::state::{ChannelFlow, Flow, CHANNEL_FLOWS, GOVMODULE, IBCMODULE};
use crate::ContractError;
use cosmwasm_std::{Addr, DepsMut, Response, Timestamp};

pub fn add_new_channels(
    deps: DepsMut,
    channels: Vec<Channel>,
    now: Timestamp,
) -> Result<(), ContractError> {
    for channel in channels {
        CHANNEL_FLOWS.save(
            deps.storage,
            &channel.name,
            &channel
                .quotas
                .iter()
                .map(|q| ChannelFlow {
                    quota: q.into(),
                    flow: Flow::new(0_u128, 0_u128, now, q.duration),
                })
                .collect(),
        )?
    }
    Ok(())
}

pub fn try_add_channel(
    deps: DepsMut,
    sender: Addr,
    channel_id: String,
    quotas: Vec<QuotaMsg>,
    now: Timestamp,
) -> Result<Response, ContractError> {
    let ibc_module = IBCMODULE.load(deps.storage)?;
    let gov_module = GOVMODULE.load(deps.storage)?;
    if sender != ibc_module && sender != gov_module {
        return Err(ContractError::Unauthorized {});
    }
    add_new_channels(
        deps,
        vec![Channel {
            name: channel_id.to_string(),
            quotas,
        }],
        now,
    )?;

    Ok(Response::new()
        .add_attribute("method", "try_add_channel")
        .add_attribute("channel_id", channel_id))
}

pub fn try_remove_channel(
    deps: DepsMut,
    sender: Addr,
    channel_id: String,
) -> Result<Response, ContractError> {
    let ibc_module = IBCMODULE.load(deps.storage)?;
    let gov_module = GOVMODULE.load(deps.storage)?;
    if sender != ibc_module && sender != gov_module {
        return Err(ContractError::Unauthorized {});
    }
    CHANNEL_FLOWS.remove(deps.storage, &channel_id);
    Ok(Response::new()
        .add_attribute("method", "try_remove_channel")
        .add_attribute("channel_id", channel_id))
}

pub fn try_reset_channel_quota(
    deps: DepsMut,
    sender: Addr,
    channel_id: String,
    quota_id: String,
    now: Timestamp,
) -> Result<Response, ContractError> {
    let gov_module = GOVMODULE.load(deps.storage)?;
    if sender != gov_module {
        return Err(ContractError::Unauthorized {});
    }

    CHANNEL_FLOWS.update(
        deps.storage,
        &channel_id.clone(),
        |maybe_flows| match maybe_flows {
            None => Err(ContractError::QuotaNotFound {
                quota_id,
                channel_id: channel_id.clone(),
            }),
            Some(mut flows) => {
                flows.iter_mut().for_each(|channel| {
                    if channel.quota.name == channel_id.as_ref() {
                        channel.flow.expire(now, channel.quota.duration)
                    }
                });
                Ok(flows)
            }
        },
    )?;

    Ok(Response::new()
        .add_attribute("method", "try_reset_channel")
        .add_attribute("channel_id", channel_id))
}

#[cfg(test)]
mod tests {

    use cosmwasm_std::testing::{mock_dependencies, mock_env, mock_info};
    use cosmwasm_std::{from_binary, Addr, StdError};

    use crate::contract::{execute, query};
    use crate::helpers::tests::verify_query_response;
    use crate::msg::{ExecuteMsg, QueryMsg, QuotaMsg};
    use crate::state::{ChannelFlow, GOVMODULE, IBCMODULE};

    const IBC_ADDR: &str = "IBC_MODULE";
    const GOV_ADDR: &str = "GOV_MODULE";

    #[test]
    fn management_add_and_remove_channel() {
        let mut deps = mock_dependencies();
        IBCMODULE
            .save(deps.as_mut().storage, &Addr::unchecked(IBC_ADDR))
            .unwrap();
        GOVMODULE
            .save(deps.as_mut().storage, &Addr::unchecked(GOV_ADDR))
            .unwrap();

        let msg = ExecuteMsg::AddChannel {
            channel_id: "channel".to_string(),
            quotas: vec![QuotaMsg {
                name: "daily".to_string(),
                duration: 1600,
                send_recv: (3, 5),
            }],
        };
        let info = mock_info(IBC_ADDR, &vec![]);

        let env = mock_env();
        let res = execute(deps.as_mut(), env.clone(), info, msg).unwrap();
        assert_eq!(0, res.messages.len());

        let query_msg = QueryMsg::GetQuotas {
            channel_id: "channel".to_string(),
        };

        let res = query(deps.as_ref(), mock_env(), query_msg.clone()).unwrap();

        let value: Vec<ChannelFlow> = from_binary(&res).unwrap();
        verify_query_response(
            &value[0],
            "daily",
            (3, 5),
            1600,
            0,
            0,
            env.block.time.plus_seconds(1600),
        );

        assert_eq!(value.len(), 1);

        // Add another channel
        let msg = ExecuteMsg::AddChannel {
            channel_id: "channel2".to_string(),
            quotas: vec![QuotaMsg {
                name: "daily".to_string(),
                duration: 1600,
                send_recv: (3, 5),
            }],
        };
        let info = mock_info(IBC_ADDR, &vec![]);

        let env = mock_env();
        execute(deps.as_mut(), env.clone(), info, msg).unwrap();

        // remove the first one
        let msg = ExecuteMsg::RemoveChannel {
            channel_id: "channel".to_string(),
        };

        let info = mock_info(IBC_ADDR, &vec![]);
        let env = mock_env();
        execute(deps.as_mut(), env.clone(), info, msg).unwrap();

        // The channel is not there anymore
        let err = query(deps.as_ref(), mock_env(), query_msg.clone()).unwrap_err();
        assert!(matches!(err, StdError::NotFound { .. }));

        // The second channel is still there
        let query_msg = QueryMsg::GetQuotas {
            channel_id: "channel2".to_string(),
        };
        let res = query(deps.as_ref(), mock_env(), query_msg.clone()).unwrap();
        let value: Vec<ChannelFlow> = from_binary(&res).unwrap();
        assert_eq!(value.len(), 1);
    }
}
