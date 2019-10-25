import ButtonSignin from 'components/button_signin'
import { iconMap } from 'components/icon'
import ListSubscriptions from 'components/list_subscriptions'
import { NextSeo } from 'next-seo'
import React from 'react'
import { connect } from 'react-redux'
import { getIsUserSignedIn } from 'selectors/entities/users'
import { AppState } from 'store'
import { SET_CURRENT_URL_PATH } from 'types/actions'
import { PageContext } from 'types/utilities'
import * as gtag from 'utils/gtag'

interface StateToProps {
  isUserSignedIn: boolean
}

interface OwnProps {
  scrollY: number
}

interface Props extends StateToProps, OwnProps {}

class SubscriptionsPage extends React.Component<Props> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    ctx.store.dispatch({
      type: SET_CURRENT_URL_PATH,
      urlPath: '/subscriptions',
    })
  }

  componentDidMount() {
    gtag.pageview('/subscriptions')
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { isUserSignedIn } = this.props

    if (!isUserSignedIn) {
      const SubscribeIcon = iconMap['history']
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
        <ListSubscriptions />
      </>
    )
  }
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    isUserSignedIn: getIsUserSignedIn(state),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  SubscriptionsPage,
)
