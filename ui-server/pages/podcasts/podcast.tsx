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
  scrollY: number
}

interface Props extends StateToProps, OwnProps {}

export default class PodcastPage extends Component<Props> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { query, store, isServer } = ctx
    const loadPodcast = bindActionCreators(getPodcast, store.dispatch)(query[
      'podcastId'
    ] as string)

    if (isServer) {
      // Type definitions for bindActionCreators is not overloaded for thunks
      // - https://github.com/piotrwitek/react-redux-typescript-guide/issues/110
      // - https://github.com/piotrwitek/react-redux-typescript-guide/issues/6
      // - https://github.com/piotrwitek/react-redux-typescript-guide/pull/157
      await loadPodcast
    }
  }

  componentDidMount() {
    window.window.scrollTo(0, this.props.scrollY)
  }

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
