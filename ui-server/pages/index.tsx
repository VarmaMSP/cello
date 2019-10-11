import { getTrendingPodcasts } from 'actions/podcast'
import Discover from 'components/discover'
import LoadingPage from 'components/loading_page'
import { NextSeo } from 'next-seo'
import React from 'react'
import { connect } from 'react-redux'
import { RequestState } from 'reducers/requests/utils'
import { bindActionCreators, Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import { logPageView } from 'utils/analytics'

interface StateToProps {
  reqState: RequestState
}

interface DispatchToProps {
  loadTrendingPodcasts: () => void
}

interface OwnProps {
  scrollY: number
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

class IndexPage extends React.Component<Props> {
  componentDidMount() {
    logPageView()
    this.props.loadTrendingPodcasts()

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
            title="Phenopod"
            description="Podcast Player for Web"
            canonical="https://phenopod.com"
            openGraph={{
              url: 'https://phenopod.com',
              type: 'website',
              title: 'Phenopod',
              description: 'Podcast Player for Web',
              site_name: 'Phenopod',
            }}
            twitter={{
              cardType: 'summary',
              site: '@phenopod',
              handle: '@phenopod',
            }}
            facebook={{
              appId: '526472207897979',
            }}
          />
          <Discover />
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

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    loadTrendingPodcasts: bindActionCreators(getTrendingPodcasts, dispatch),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(IndexPage)
