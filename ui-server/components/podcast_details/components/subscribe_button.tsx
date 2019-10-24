import { subscribeToPodcast, unsubscribeToPodcast } from 'actions/podcast'
import classNames from 'classnames'
import React from 'react'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getIsCurrentUserSubscribedToPodcast } from 'selectors/entities/podcasts'
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

const SubscribeButton: React.SFC<Props> = (props) => {
  const { isSubscribed, unsubscribe, subscribe, podcastId } = props
  return (
    <button
      className={classNames(
        props.className,
        'rounded tracking-tight focus:outline-none focus:shadow-outline',
        {
          'bg-indigo-500 text-white': !isSubscribed,
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
    isSubscribed: getIsCurrentUserSubscribedToPodcast(state, podcastId),
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
)(SubscribeButton)
