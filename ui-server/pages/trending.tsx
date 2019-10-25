import { getTrendingPodcasts } from 'actions/podcast'
import ListTrendingPodcasts from 'components/list_trending_podcasts'
import LoadingPage from 'components/loading_page'
import React from 'react'
import { connect } from 'react-redux'
import { RequestState } from 'reducers/requests/utils'
import { bindActionCreators } from 'redux'
import { AppState } from 'store'
import { PageContext } from 'types/utilities'
import * as gtag from 'utils/gtag'

interface StateToProps {
  reqState: RequestState
}

interface OwnProps {
  scrollY: number
}

class TrendingPage extends React.Component<StateToProps & OwnProps> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { store } = ctx
    await bindActionCreators(getTrendingPodcasts, store.dispatch)()
  }

  componentDidMount() {
    gtag.pageview('/trending')

    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { reqState } = this.props

    if (reqState.status == 'STARTED') {
      return <LoadingPage />
    }
    if (reqState.status == 'SUCCESS') {
      return <ListTrendingPodcasts />
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
