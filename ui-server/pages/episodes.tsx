import { getEpisode } from 'actions/episode'
import EpisodeView from 'components/episode_view'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { PageContext } from 'types/utilities'

interface OwnProps {
  episodeId: string
  scrollY: number
}

export default class EpisodePage extends Component<OwnProps> {
  static async getInitialProps({ query, store }: PageContext): Promise<void> {
    if (!!query['skipLoad']) {
      return
    }

    await bindActionCreators(
      getEpisode,
      store.dispatch,
    )(query['episodeId'] as string)
  }

  componentDidMount() {
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    return <EpisodeView episodeId={this.props.episodeId} />
  }
}
