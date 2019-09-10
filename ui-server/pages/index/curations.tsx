import LoadingPage from 'components/loading_page'
import React from 'react'
import { RequestState } from 'reducers/requests/utils'
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
  static async getInitialProps(_ctx: PageContext) {
    return { preventInitalLoad: false }
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
      return <>{'curations loaded'}</>
    }
    return <>Hey Morty</>
  }
}
