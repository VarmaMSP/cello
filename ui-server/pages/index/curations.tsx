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

export interface DispatchToProps {
  loadCurations: () => void
}

interface Props extends StateToProps, DispatchToProps {}

export default class extends React.Component<Props> {
  static async getInitialProps({ store }: PageContext): Promise<{}> {
    const loadCurations = bindActionCreators(getCurations, store.dispatch)

    await loadCurations()
    return {}
  }

  render() {
    const { reqState } = this.props

    if (reqState.status == 'STARTED' || reqState.status == 'NOT_STARTED') {
      return <LoadingPage />
    }

    if (reqState.status == 'SUCCESS') {
      return <Discover />
    }
    return <>Hey Morty</>
  }
}