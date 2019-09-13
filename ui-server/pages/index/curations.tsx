import { getCurations } from 'actions/curations'
import Discover from 'components/discover'
import LoadingPage from 'components/loading_page'
import React from 'react'
import { RequestState } from 'reducers/requests/utils'
import { bindActionCreators } from 'redux'
import { PageContext } from 'types/utilities'

export interface StateToProps {
  reqState: RequestState
}

export default class extends React.Component<StateToProps> {
  static async getInitialProps(ctx: PageContext): Promise<{}> {
    const { store, isServer } = ctx
    const loadCurations = bindActionCreators(getCurations, store.dispatch)()

    if (isServer) {
      await loadCurations
    }
    return {}
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
