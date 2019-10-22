import { getTrendingPodcasts } from 'actions/podcast'
import Discover from 'components/discover'
import LoadingPage from 'components/loading_page'
import { NextSeo } from 'next-seo'
import React from 'react'
import { connect } from 'react-redux'
import { RequestState } from 'reducers/requests/utils'
import { bindActionCreators, Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions, SET_CURRENT_URL_PATH } from 'types/actions'
import { PageContext } from 'types/utilities'
import * as gtag from 'utils/gtag'

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
  static async getInitialProps(ctx: PageContext): Promise<void> {
    ctx.store.dispatch({ type: SET_CURRENT_URL_PATH, urlPath: '/' })
  }

  componentDidMount() {
    gtag.pageview('/')

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
