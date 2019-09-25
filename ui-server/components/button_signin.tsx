import React from 'react'
import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions, SHOW_SIGNIN_MODAL } from 'types/actions'

interface DispatchToProps {
  showSigninModal: () => void
}

const ButtonSignin: React.SFC<DispatchToProps> = (props) => {
  return (
    <button
      className="w-full h-full rounded border-2 border-orange-600 focus:outline-none focus:shadow-outline"
      onClick={props.showSigninModal}
    >
      <p className="text-sm text-orange-600 font-semibold leading-loose">
        SIGN IN
      </p>
    </button>
  )
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    showSigninModal: () => dispatch({ type: SHOW_SIGNIN_MODAL }),
  }
}

export default connect<{}, DispatchToProps, {}, AppState>(
  null,
  mapDispatchToProps,
)(ButtonSignin)
