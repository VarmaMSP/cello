import { getSubscriptionsPageData } from 'actions/subscription'
import ButtonSignin from 'components/button_signin'
import { iconMap } from 'components/icon'
import PageLayout from 'components/page_layout'
import SubscriptionsFeed from 'components/subscriptions_feed'
import SubscriptionsList from 'components/subscriptions_list'
import { NextSeo } from 'next-seo'
import React from 'react'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getSubscriptionsPageStatus } from 'selectors/request'
import { getIsUserSignedIn } from 'selectors/session'
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
    gtag.pageview('/subscriptions')
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { isUserSignedIn, isLoading } = this.props

    if (!isUserSignedIn || isLoading) {
      const SubscribeIcon = iconMap['heart']
      return (
        <>
          <div className="mx-auto mt-24">
            <SubscribeIcon className="w-12 h-12 mx-auto fill-current text-gray-700" />
            <h1 className="text-center text-xl text-gray-700 my-6">
              {'Sign in to subscribe to your favourite podcasts'}
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
          title="Subscriptions - Phenopod"
          description="Subscriptions"
          canonical="https://phenopod.com/subscripitions"
        />
        <PageLayout>
          <SubscriptionsFeed />
          <SubscriptionsList />
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
