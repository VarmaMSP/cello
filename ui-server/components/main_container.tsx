import React, { Component } from 'react'
import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions, SET_SCREEN_WIDTH } from 'types/actions'
import { ScreenWidth } from 'types/app'

interface DispatchToProps {
  setScreenWidth: (s: ScreenWidth) => void
}

interface OwnProps {
  children: JSX.Element | JSX.Element[]
}

interface Props extends DispatchToProps, OwnProps {}

class Screen extends Component<Props> {
  componentDidMount() {
    window.addEventListener('resize', this.handleScreenResize)
    this.handleScreenResize()
  }

  componentWillUnmount() {
    window.removeEventListener('resize', () => {})
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
      <div className="lg:pl-56 pl-4 lg:pr-5 pr-4 pt-20 pb-64 z-0">
        <div className="lg:pl-5 lg:pb-36">{children}</div>
      </div>
    )
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    setScreenWidth: (s: ScreenWidth) =>
      dispatch({ type: SET_SCREEN_WIDTH, width: s }),
  }
}

export default connect<{}, DispatchToProps, {}, AppState>(
  null,
  mapDispatchToProps,
)(Screen)
