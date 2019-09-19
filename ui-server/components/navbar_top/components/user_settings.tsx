import React, { useEffect, useState } from 'react'
import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { getCurrenUser } from 'selectors/entities/users'
import { AppState } from 'store'
import { AppActions, USER_SIGNED_OUT } from 'types/actions'
import { User } from 'types/app'

interface StateToProps {
  user: User
}

interface DispatchToProps {
  signOutUser: () => void
}

const UserSettings: React.SFC<StateToProps & DispatchToProps> = (props) => {
  const dropdown = React.createRef<HTMLDivElement>()
  const { user, signOutUser } = props
  const [showDropDown, setShowDropdown] = useState(false)

  useEffect(() => {
    document.addEventListener('mousedown', (e) => {
      if (dropdown.current && !dropdown.current.contains(e.target as any)) {
        setShowDropdown(false)
      }
    })
  })

  return (
    <div className="relative w-full h-full">
      <button
        className="w-full h-full rounded border-2 border-blue-600 focus:outline-none focus:shadow-outline"
        onClick={() => setShowDropdown(!showDropDown)}
      >
        <p className="text-sm text-blue-600 font-semibold leading-loose">
          {user.name}
        </p>
      </button>
      {showDropDown && (
        <div
          className="absolute right-0 w-36 z-50 py-5 border-gray-300 shadow-lg rounded"
          style={{ top: '110%' }}
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
    signOutUser: () => dispatch({ type: USER_SIGNED_OUT }),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(UserSettings)
