import React from 'react'
import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions, MODAL_MANAGER_SHOW_SIGN_IN_MODAL } from 'types/actions'

interface DispatchToProps {
  showSigninModal: () => void
}

const ButtonSignin: React.SFC<DispatchToProps> = (props) => {
  return (
    <button
      className="w-full h-full rounded bg-orange-600 focus:outline-none focus:shadow-outline"
      onClick={props.showSigninModal}
    >
      <p className="text-sm text-white font-semibold leading-loose">SIGN IN</p>
    </button>
  )
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    showSigninModal: () => dispatch({ type: MODAL_MANAGER_SHOW_SIGN_IN_MODAL }),
  }
}

export default connect<{}, DispatchToProps, {}, AppState>(
  null,
  mapDispatchToProps,
)(ButtonSignin)
