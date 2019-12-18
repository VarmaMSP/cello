import { getEpisodePageData } from 'actions/episode'
import EpisodeView from 'components/episode_view'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { PageContext } from 'types/utilities'
import { getIdFromUrlParam } from 'utils/format'

interface OwnProps {
  episodeUrlParam: string
  scrollY: number
}

export default class EpisodePage extends Component<OwnProps> {
  static async getInitialProps({ query, store }: PageContext): Promise<void> {
    if (!!query['skipLoad']) {
      return
    }

    await bindActionCreators(
      getEpisodePageData,
      store.dispatch,
    )(query['episodeUrlParam'] as string)
  }

  componentDidMount() {
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    return (
      <EpisodeView episodeId={getIdFromUrlParam(this.props.episodeUrlParam)} />
    )
  }
}
