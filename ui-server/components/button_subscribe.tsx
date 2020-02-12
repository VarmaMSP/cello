import { subscribeToPodcast, unsubscribeToPodcast } from 'actions/subscription'
import classNames from 'classnames'
import React from 'react'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getIsSubscribed } from 'selectors/session'
import { AppState } from 'store'
import { AppActions } from 'types/actions'

interface StateToProps {
  isSubscribed: boolean
}

interface DispatchToProps {
  subscribe: (podcastId: string) => void
  unsubscribe: (podcastId: string) => void
}

interface OwnProps {
  className: string
  podcastId: string
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

const ButtonSubscribe: React.SFC<Props> = (props) => {
  const { isSubscribed, unsubscribe, subscribe, podcastId } = props
  return (
    <button
      className={classNames(
        props.className,
        'rounded font-semibold tracking-wide focus:outline-none focus:shadow-outline',
        {
          'bg-purple-600 text-white': !isSubscribed,
          'bg-gray-300 text-gray-700': isSubscribed,
        },
      )}
      onClick={() =>
        isSubscribed ? unsubscribe(podcastId) : subscribe(podcastId)
      }
    >
      {!isSubscribed ? 'SUBSCRIBE' : 'SUBSCRIBED'}
    </button>
  )
}

function mapStateToProps(
  state: AppState,
  { podcastId }: OwnProps,
): StateToProps {
  return {
    isSubscribed: getIsSubscribed(state, podcastId),
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    subscribe: bindActionCreators(subscribeToPodcast, dispatch),
    unsubscribe: bindActionCreators(unsubscribeToPodcast, dispatch),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(ButtonSubscribe)
