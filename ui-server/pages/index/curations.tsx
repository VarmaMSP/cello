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

export interface OwnProps {
  preventInitialLoad: boolean
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

export default class extends React.Component<Props> {
  static async getInitialProps({
    isServer,
    store,
  }: PageContext): Promise<OwnProps> {
    const loadCurations = bindActionCreators(getCurations, store.dispatch)

    await loadCurations()
    return { preventInitialLoad: isServer }
  }

  componentDidMount() {
    if (!this.props.preventInitialLoad) {
      this.props.loadCurations()
    }
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
