import { getEpisodePageData } from 'actions/episode'
import EpisodeView from 'components/episode_view'
import PageLayout from 'components/page_layout'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { PageContext } from 'types/utilities'
import { getIdFromUrlParam } from 'utils/format'

interface OwnProps {
  episodeUrlParam: string
  activeTab?: string
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
    const { episodeUrlParam } = this.props
    const episodeId = getIdFromUrlParam(episodeUrlParam)

    return (
      <PageLayout>
        <EpisodeView episodeId={episodeId} />
        <div />
      </PageLayout>
    )
  }
}
