import { getSubscriptionsPageData } from 'actions/subscription'
import { iconMap } from 'components/icon'
import PageLayout from 'components/page_layout'
import { SubscriptionsPageSeo } from 'components/seo'
import SignInButton from 'components/sign_in_button'
import SubscriptionsFeed from 'components/subscriptions_feed'
import SubscriptionsList from 'components/subscriptions_list'
import React from 'react'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getSubscriptionsPageStatus } from 'selectors/request'
import { getIsUserSignedIn } from 'selectors/session'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import { PageContext } from 'types/utilities'

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

type Props = StateToProps & DispatchToProps & OwnProps

class SubscriptionsPage extends React.Component<Props> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { store } = ctx

    if (getIsUserSignedIn(store.getState())) {
      await bindActionCreators(getSubscriptionsPageData, store.dispatch)()
    }
  }

  componentDidUpdate(prevProps: Props) {
    const { isUserSignedIn } = this.props
    if (isUserSignedIn && !prevProps.isUserSignedIn) {
      this.props.loadPageData()
    }
  }

  componentDidMount() {
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { isUserSignedIn, isLoading } = this.props

    if (!isUserSignedIn || isLoading) {
      const SubscribeIcon = iconMap['heart']
      return (
        <PageLayout>
          <div className="mx-auto mt-32">
            <SubscribeIcon className="w-12 h-12 mx-auto fill-current text-gray-700" />
            <h1 className="text-center text-xl text-gray-700 my-6">
              {'Sign in to subscribe to your favourite podcasts'}
            </h1>
            <div className="h-10 w-48 mx-auto">
              <SignInButton />
            </div>
          </div>
        </PageLayout>
      )
    }

    return (
      <>
        <SubscriptionsPageSeo />
        <PageLayout>
          <SubscriptionsFeed />
          <div className="pt-4">
            <SubscriptionsList />
          </div>
        </PageLayout>
      </>
    )
  }
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    isUserSignedIn: getIsUserSignedIn(state),
    isLoading: getSubscriptionsPageStatus(state) === 'IN_PROGRESS',
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    loadPageData: bindActionCreators(getSubscriptionsPageData, dispatch),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(SubscriptionsPage)
