import { getTrendingPodcasts } from 'actions/podcast'
import ListTrendingPodcasts from 'components/list_trending_podcasts'
import LoadingPage from 'components/loading_page'
import { NextSeo } from 'next-seo'
import React from 'react'
import { connect } from 'react-redux'
import { RequestState } from 'reducers/requests/utils'
import { bindActionCreators } from 'redux'
import { AppState } from 'store'
import { PageContext } from 'types/utilities'
import { logPageView } from 'utils/analytics'

interface StateToProps {
  reqState: RequestState
}

interface OwnProps {
  scrollY: number
}

class TrendingPage extends React.Component<StateToProps & OwnProps> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { store, isServer } = ctx
    const loadTrendingPodcasts = bindActionCreators(
      getTrendingPodcasts,
      store.dispatch,
    )()

    if (isServer) {
      await loadTrendingPodcasts
    }
  }

  componentDidMount() {
    logPageView()

    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { reqState } = this.props

    if (reqState.status == 'STARTED') {
      return <LoadingPage />
    }
    if (reqState.status == 'SUCCESS') {
      return (
        <>
          <NextSeo
            title="Trending Podcasts - Phenopod"
            description="Trending podcasts"
            canonical="https://phenopod.com/trending"
            openGraph={{
              url: 'https://phenopod.com/trending',
              type: 'article',
              title: 'Trending Podcasts',
              description: 'Trending Podcasts',
              site_name: 'Phenopod',
            }}
            twitter={{
              cardType: `summary`,
              site: '@phenopod',
              handle: '@phenopod',
            }}
            facebook={{
              appId: '526472207897979',
            }}
          />
          <ListTrendingPodcasts />
        </>
      )
    }
    return <></>
  }
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    reqState: state.requests.podcast.getTrendingPodcasts,
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  TrendingPage,
)
