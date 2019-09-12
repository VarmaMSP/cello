import React from 'react'
import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions, SHOW_SIGN_IN_MODAL } from 'types/actions'

interface DispatchToProps {
  showModal: () => void
}

const SignInButton: React.SFC<DispatchToProps> = (props) => {
  return (
    <button
      className="md:w-24 w-20 h-8 bg-orange-600 hover:bg-orange-700 active:bg-orange-700 rounded-lg focus:outline-none focus:shadow-outline"
      onClick={props.showModal}
    >
      <p className="text-sm text-white font-semibold leading-loose">SIGN IN</p>
    </button>
  )
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    showModal: () => dispatch({ type: SHOW_SIGN_IN_MODAL }),
  }
}

export default connect<{}, DispatchToProps, {}, AppState>(
  null,
  mapDispatchToProps,
)(SignInButton)
