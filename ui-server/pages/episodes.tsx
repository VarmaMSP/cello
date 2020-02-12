import { getEpisodePageData } from 'actions/episode'
import EpisodeView from 'components/episode_view'
import PageLayout from 'components/page_layout'
import { EpisodePageSeo } from 'components/seo'
import React, { Component } from 'react'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import { getEpisodeById } from 'selectors/entities/episodes'
import { AppState } from 'store'
import { Episode } from 'models'
import { PageContext } from 'types/utilities'
import { getIdFromUrlParam } from 'utils/format'

export interface StateToProps { 
  episode: Episode
}

interface OwnProps {
  episodeUrlParam: string
  activeTab?: string
  scrollY: number
}

class EpisodePage extends Component<StateToProps & OwnProps> {
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
    const { episode, activeTab } = this.props

    return (
      <>
        <EpisodePageSeo episode={episode} />
        <PageLayout>
          <EpisodeView episode={episode} activeTab={activeTab} />
          <div />
        </PageLayout>
      </>
    )
  }
}

function mapStateToProps(
  state: AppState,
  { episodeUrlParam }: OwnProps,
): StateToProps {
  return { episode: getEpisodeById(state, getIdFromUrlParam(episodeUrlParam)) }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  EpisodePage,
)
