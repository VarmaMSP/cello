import { getPodcast } from 'actions/podcast'
import EpisodeList from 'components/episode_list'
import LoadingPage from 'components/loading_page'
import PodcastDetails from 'components/podcast_details'
import React, { Component } from 'react'
import { RequestState } from 'reducers/requests/utils'
import { bindActionCreators } from 'redux'
import { PageContext } from 'types/utilities'

export interface StateToProps {
  reqState: RequestState
}

export interface OwnProps {
  podcastId: string
}

interface Props extends StateToProps, OwnProps {}

export default class PodcastPage extends Component<Props> {
  static async getInitialProps(ctx: PageContext): Promise<OwnProps> {
    const { query, store, isServer } = ctx
    const podcastId = query['id'] as string
    const loadPodcast = bindActionCreators(getPodcast, store.dispatch)(
      podcastId,
    )

    if (isServer) {
      // Type definitions for bindActionCreators is not overloaded for thunks
      // - https://github.com/piotrwitek/react-redux-typescript-guide/issues/110
      // - https://github.com/piotrwitek/react-redux-typescript-guide/issues/6
      // - https://github.com/piotrwitek/react-redux-typescript-guide/pull/157
      await loadPodcast
    }
    return { podcastId }
  }

  // Calling loadPodcast in componentDidUpdate is not needed, a new url is pushed into the history for
  // every new page so it will be called in the getInitalProps itself

  render() {
    const { reqState, podcastId } = this.props

    if (reqState.status == 'STARTED') {
      return <LoadingPage />
    }

    if (reqState.status == 'SUCCESS') {
      return (
        <>
          <PodcastDetails podcastId={podcastId} />
          <EpisodeList podcastId={podcastId} />
        </>
      )
    }

    return <></>
  }
}
