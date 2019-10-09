import { getTrendingPodcasts } from 'actions/podcast'
import Discover from 'components/discover'
import LoadingPage from 'components/loading_page'
import React from 'react'
import { connect } from 'react-redux'
import { RequestState } from 'reducers/requests/utils'
import { bindActionCreators } from 'redux'
import { AppState } from 'store'
import { PageContext } from 'types/utilities'

interface StateToProps {
  reqState: RequestState
}

interface OwnProps {
  scrollY: number
}

class IndexPage extends React.Component<StateToProps & OwnProps> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { store, isServer } = ctx
    const loadTrendingPodcasts = bindActionCreators(
      getTrendingPodcasts,
      store.dispatch,
    )()

    if (!isServer) {
      await loadTrendingPodcasts
    }
  }

  componentDidMount() {
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { reqState } = this.props

    if (reqState.status == 'STARTED') {
      return <LoadingPage />
    }
    if (reqState.status == 'SUCCESS') {
      return <Discover />
    }
    return <>Hey Morty</>
  }
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    reqState: state.requests.podcast.getTrendingPodcasts,
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  IndexPage,
)
