import React, { Component } from 'react'
import { ScreenWidth } from '../types/app'
import { Dispatch } from 'redux'
import { AppActions, SET_SCREEN_WIDTH } from '../types/actions'
import { connect } from 'react-redux'
import { AppState } from 'store'

interface DispatchToProps {
  setScreenWidth: (s: ScreenWidth) => void
}

class Screen extends Component<DispatchToProps> {
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
    return <></>
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
