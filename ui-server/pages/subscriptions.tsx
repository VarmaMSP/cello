import ListSubscriptions from 'components/list_subscriptions'
import LoadingPage from 'components/loading_page'
import { NextSeo } from 'next-seo'
import React from 'react'
import { connect } from 'react-redux'
import { RequestState } from 'reducers/requests/utils'
import { AppState } from 'store'
import { logPageView } from 'utils/analytics'

interface StateToProps {
  reqState: RequestState
}

interface OwnProps {
  scrollY: number
}

class SubscriptionsPage extends React.Component<StateToProps & OwnProps> {
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
            noindex
            title="Subscriptions - Phenopod"
            description="Subscriptions"
            canonical="https://phenopod.com/subscripitions"
          />
          <ListSubscriptions />
        </>
      )
    }
    return <></>
  }
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    reqState: state.requests.user.getSignedInUser,
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  SubscriptionsPage,
)
