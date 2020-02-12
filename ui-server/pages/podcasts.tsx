import { getPodcastPageData } from 'actions/podcast'
import PageLayout from 'components/page_layout'
import PodcastView from 'components/podcast_view'
import { PodcastPageSeo } from 'components/seo'
import React, { Component } from 'react'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import { getPodcastById } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import { Podcast } from 'models'
import { PageContext } from 'types/utilities'
import { getIdFromUrlParam } from 'utils/utils'

interface StateToProps {
  podcast: Podcast
}

interface OwnProps {
  podcastUrlParam: string
  activeTab?: string
  scrollY: number
}

class PodcastPage extends Component<StateToProps & OwnProps> {
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
    const { podcast, activeTab } = this.props

    return (
      <>
        <PodcastPageSeo podcast={podcast} />
        <PageLayout>
          <PodcastView podcast={podcast} activeTab={activeTab} />
          <div></div>
        </PageLayout>
      </>
    )
  }
}

function mapStateToProps(
  state: AppState,
  { podcastUrlParam }: OwnProps,
): StateToProps {
  return {
    podcast: getPodcastById(state, getIdFromUrlParam(podcastUrlParam)),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  PodcastPage,
)
