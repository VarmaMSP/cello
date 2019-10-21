import { getUserFeed } from 'actions/user'
import ListFeed from 'components/list_feed'
import LoadingPage from 'components/loading_page'
import { NextSeo } from 'next-seo'
import React from 'react'
import { connect } from 'react-redux'
import { RequestState } from 'reducers/requests/utils'
import { bindActionCreators, Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions, SET_CURRENT_URL_PATH } from 'types/actions'
import { PageContext } from 'types/utilities'
import { logPageView } from 'utils/analytics'

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
  static async getInitialProps(ctx: PageContext): Promise<void> {
    ctx.store.dispatch({ type: SET_CURRENT_URL_PATH, urlPath: '/feed' })
  }

  componentDidMount() {
    logPageView()
    this.props.loadFeed()

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
            noindex
            title="Feed - Phenopod"
            description="Feed"
            canonical="https://phenopod.com/feed"
          />
          <ListFeed />
        </>
      )
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
