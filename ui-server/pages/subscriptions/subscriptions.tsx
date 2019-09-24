import LoadingPage from 'components/loading_page'
import SubscriptionsList from 'components/subscriptions_list'
import React from 'react'
import { RequestState } from 'reducers/requests/utils'

export interface StateToProps {
  reqState: RequestState
}

export interface OwnProps {
  scrollY: number
}

interface Props extends StateToProps, OwnProps {}

export default class extends React.Component<Props> {
  componentDidMount() {
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { reqState } = this.props

    if (reqState.status == 'STARTED') {
      return <LoadingPage />
    }

    if (reqState.status == 'SUCCESS') {
      return <SubscriptionsList />
    }

    return <></>
  }
}
