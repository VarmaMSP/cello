import classnames from 'classnames'
import React from 'react'
import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions, MODAL_MANAGER_SHOW_SIGN_IN_MODAL } from 'types/actions'

interface DispatchToProps {
  showSignInModal: () => void
}

interface OwnProps {
  small?: boolean
}

const ButtonSignin: React.SFC<DispatchToProps & OwnProps> = ({
  showSignInModal,
  small,
}) => {
  return (
    <button
      className={classnames(
        'w-full h-full bg-orange-600 text-gray-100 leading-loose',
        'tracking-wide focus:outline-none focus:shadow-outline',
        {
          'rounded-full': !small,
          'rounded-lg': small,
        },
      )}
      onClick={showSignInModal}
    >
      {small ? (
        <p className="text-sm">{'SIGN IN'}</p>
      ) : (
        <p className="text-sm font-medium">{'SIGN IN / SIGN UP'}</p>
      )}
    </button>
  )
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    showSignInModal: () => dispatch({ type: MODAL_MANAGER_SHOW_SIGN_IN_MODAL }),
  }
}

export default connect<{}, DispatchToProps, {}, AppState>(
  null,
  mapDispatchToProps,
)(ButtonSignin)
