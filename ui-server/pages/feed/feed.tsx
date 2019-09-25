import FeedList from 'components/feed_list'
import LoadingPage from 'components/loading_page'
import React from 'react'
import { RequestState } from 'reducers/requests/utils'

export interface StateToProps {
  reqState: RequestState
}

export interface DispatchToProps {
  loadFeed: () => void
}

export interface OwnProps {
  scrollY: number
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

export default class extends React.Component<Props> {
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
      return <FeedList />
    }

    return <></>
  }
}
