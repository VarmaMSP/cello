import { getSignedInUser } from 'actions/user'
import Router from 'next/router'
import React, { Component } from 'react'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { AppState } from 'store'
import {
  AppActions,
  SET_CURRENT_PATH_NAME,
  SET_SCREEN_WIDTH,
} from 'types/actions'
import { ScreenWidth } from 'types/app'

interface DispatchToProps {
  getSignedInUser: () => void
  setScreenWidth: (s: ScreenWidth) => void
  setCurrentPathName: (s: string) => void
}

interface OwnProps {
  children: JSX.Element | JSX.Element[]
}

interface Props extends DispatchToProps, OwnProps {}

class Screen extends Component<Props> {
  componentDidMount() {
    this.props.getSignedInUser()
    this.props.setCurrentPathName(Router.pathname)
    this.handleScreenResize()

    window.addEventListener('resize', this.handleScreenResize)
    Router.events.on('routeChangeStart', this.props.setCurrentPathName)
  }

  componentWillUnmount() {
    window.removeEventListener('resize', () => {})
    Router.events.off('routeChangeStart', () => {})
  }

  handleScreenResize = () => {
    const screenWidth = window.innerWidth
    if (screenWidth >= 1024) return this.props.setScreenWidth('LG')
    if (screenWidth >= 768) return this.props.setScreenWidth('MD')
    this.props.setScreenWidth('SM')
  }

  render() {
    const { children } = this.props

    return (
      // base padding
      <div className="pl-4 pr-4 pt-20 pb-64 z-0">
        {/* additonal padding for large screens */}
        <div className="lg:pl-60 lg:pr-1 lg:pb-36">
          {/* additonal padding for extra large screens */}
          <div className="xl:pl-20 xl:pr-40">{children}</div>
        </div>
      </div>
    )
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    getSignedInUser: bindActionCreators(getSignedInUser, dispatch),
    setScreenWidth: (s: ScreenWidth) =>
      dispatch({ type: SET_SCREEN_WIDTH, width: s }),
    setCurrentPathName: (s: string) =>
      dispatch({ type: SET_CURRENT_PATH_NAME, pathName: s }),
  }
}

export default connect<{}, DispatchToProps, {}, AppState>(
  null,
  mapDispatchToProps,
)(Screen)
