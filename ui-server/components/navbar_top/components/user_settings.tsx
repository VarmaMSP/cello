import { signOutUser } from 'actions/user'
import ButtonWithIcon from 'components/button_with_icon'
import React, { useEffect, useState } from 'react'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getCurrenUser } from 'selectors/entities/users'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import { User } from 'types/app'

interface StateToProps {
  user: User
}

interface DispatchToProps {
  signOutUser: () => void
}

const UserSettings: React.SFC<StateToProps & DispatchToProps> = (props) => {
  const dropdown = React.createRef<HTMLDivElement>()
  const [showDropDown, setShowDropdown] = useState(false)

  useEffect(() => {
    document.addEventListener('mousedown', (e) => {
      if (dropdown.current && !dropdown.current.contains(e.target as any)) {
        setShowDropdown(false)
      }
    })
  })

  const { signOutUser } = props

  return (
    <div className="relative w-full h-full">
      <ButtonWithIcon
        icon="user-solid-circle"
        className="absolute right-0 w-8 h-auto text-gray-700"
        onClick={() => setShowDropdown(!showDropDown)}
      />
      {showDropDown && (
        <div
          className="absolute right-0 w-36 z-50 py-5 bg-white border border-gray-400 shadow rounded"
          style={{ top: '130%' }}
          ref={dropdown}
        >
          <div
            className="px-4 py-2 hover:bg-gray-200 text-gray-800 cursor-pointer"
            onClick={signOutUser}
          >
            Sign Out
          </div>
        </div>
      )}
    </div>
  )
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    user: getCurrenUser()(state),
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
