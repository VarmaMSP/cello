import { getUserFeed } from 'actions/user'
import ListFeed from 'components/list_feed'
import LoadingPage from 'components/loading_page'
import React from 'react'
import { connect } from 'react-redux'
import { RequestState } from 'reducers/requests/utils'
import { bindActionCreators, Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions } from 'types/actions'

interface StateToProps {
  reqState: RequestState
}

interface DispatchToProps {
  loadFeed: () => void
}

interface OwnProps {
  scrollY: number
}

class FeedPage extends React.Component<
  StateToProps & DispatchToProps & OwnProps
> {
  static async getInitialProps(): Promise<void> {}

  componentDidMount() {
    this.props.loadFeed()
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { reqState } = this.props

    if (reqState.status == 'STARTED') {
      return <LoadingPage />
    }
    if (reqState.status == 'SUCCESS') {
      return <ListFeed />
    }
    return <></>
  }
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    reqState: state.requests.user.getUserFeed,
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    loadFeed: bindActionCreators(getUserFeed, dispatch),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(FeedPage)
