import { getHistoryPageData } from 'actions/history'
import ButtonSignin from 'components/button_signin'
import HistoryView from 'components/history_view/history_view'
import { iconMap } from 'components/icon'
import { NextSeo } from 'next-seo'
import React from 'react'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getIsUserSignedIn } from 'selectors/entities/users'
import { getHistoryPageStatus } from 'selectors/request'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import { PageContext } from 'types/utilities'
import * as gtag from 'utils/gtag'

interface StateToProps {
  isUserSignedIn: boolean
  isLoading: boolean
}

interface DispatchToProps {
  loadPageData: () => void
}

interface OwnProps {
  scrollY: number
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

class FeedPage extends React.Component<Props> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { store } = ctx

    if (getIsUserSignedIn(store.getState())) {
      await bindActionCreators(getHistoryPageData, store.dispatch)()
    }
  }

  componentDidUpdate(prevProps: Props) {
    const { isUserSignedIn } = this.props
    if (isUserSignedIn && !prevProps.isUserSignedIn) {
      this.props.loadPageData()
    }
  }

  componentDidMount() {
    gtag.pageview('/history')
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { isUserSignedIn, isLoading } = this.props

    if (!isUserSignedIn || isLoading) {
      const HistoryIcon = iconMap['history']
      return (
        <>
          <div className="mx-auto mt-24">
            <HistoryIcon className="w-12 h-12 mx-auto fill-current text-gray-700" />
            <h1 className="text-center text-xl text-gray-700 my-6">
              {'Sign in to keep track of what you listen'}
            </h1>
            <div className="w-32 mx-auto">
              <ButtonSignin />
            </div>
          </div>
        </>
      )
    }

    return (
      <>
        <NextSeo
          noindex
          title="History - Phenopod"
          description="History"
          canonical="https://phenopod.com/feed"
        />
        <HistoryView />
      </>
    )
  }
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    isUserSignedIn: getIsUserSignedIn(state),
    isLoading: getHistoryPageStatus(state) === 'IN_PROGRESS',
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    loadPageData: bindActionCreators(getHistoryPageData, dispatch),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(FeedPage)
