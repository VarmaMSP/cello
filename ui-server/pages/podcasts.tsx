import { getPodcastPageData } from 'actions/podcast'
import PageLayout from 'components/page_layout'
import PodcastView from 'components/podcast_view/podcast_view'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { PageContext } from 'types/utilities'
import { getIdFromUrlParam } from 'utils/utils'

interface OwnProps {
  podcastUrlParam: string
  activeTab?: string
  scrollY: number
}

export default class PodcastsPage extends Component<OwnProps> {
  static async getInitialProps({ query, store }: PageContext): Promise<void> {
    await bindActionCreators(
      getPodcastPageData,
      store.dispatch,
    )(query['podcastUrlParam'] as string)
  }

  componentDidMount() {
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { podcastUrlParam, activeTab } = this.props
    const podcastId = getIdFromUrlParam(podcastUrlParam)

    return (
      <PageLayout>
        <PodcastView podcastId={podcastId} activeTab={activeTab} />
        <div></div>
      </PageLayout>
    )
  }
}

