import { signOutUser } from 'actions/user'
import ButtonWithIcon from 'components/button_with_icon'
import Dropdown from 'components/dropdown'
import React from 'react'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getUserById } from 'selectors/entities/users'
import { getSignedInUserId } from 'selectors/session'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import { User } from 'types/models'

interface StateToProps {
  user: User
}

interface DispatchToProps {
  signOutUser: () => void
}

const UserSettings: React.SFC<StateToProps & DispatchToProps> = (props) => {
  const { user, signOutUser } = props

  return (
    <Dropdown
      button={
        <ButtonWithIcon
          icon="user-solid-circle"
          className="w-8 h-auto text-gray-700"
        />
      }
      dropdown={
        <div className="w-64 px-4">
          <div className="text-center">
            <div className="text-base text-gray-700">{user.name}</div>
            <div className="text-sm text-gray-700">{user.email}</div>
          </div>
          <hr className="mt-4 mb-2" />

          <button
            className="w-full py-2 hover:bg-gray-200 text-gray-800 rounded-full cursor-pointer"
            onClick={signOutUser}
          >
            Sign Out
          </button>
        </div>
      }
    />
  )
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    user: getUserById(state, getSignedInUserId(state)),
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    signOutUser: bindActionCreators(signOutUser, dispatch),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(UserSettings)
