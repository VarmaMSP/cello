import classNames from 'classnames'
import { Icon, iconMap } from 'components/icon'
import Link from 'next/link'
import React from 'react'
import { connect } from 'react-redux'
import { getCurrentUrlPath } from 'selectors/browser/urlPath'
import { getIsUserSignedIn } from 'selectors/entities/users'
import { AppState } from 'store'

interface StateToProps {
  userSignedIn: boolean
  currentUrlPath: string
}

const MenuItem: React.SFC<{ icon: Icon; name: string; active?: boolean }> = ({
  icon,
  name,
  active,
}) => {
  const Icon = iconMap[icon]
  return (
    <a>
      <div className="w-20 text-center cursor-pointer">
        <Icon
          className={classNames('w-5 h-5 mx-auto fill-current', {
            'text-gray-700': !active,
            'text-green-600': active,
          })}
        />
        <h4 className="capitalize text-sm leading-loose tracking-wide">
          {name}
        </h4>
      </div>
    </a>
  )
}

const NavbarBottom: React.SFC<StateToProps> = ({ currentUrlPath }) => {
  return (
    <div className="flex justify-around h-full pt-2 pb-0 bg-white z-50">
      <Link href="/" scroll={false}>
        <div>
          <MenuItem
            icon="explore"
            name="home"
            active={currentUrlPath === '/'}
          />
        </div>
      </Link>
      <Link href="/feed" scroll={false}>
        <div>
          <MenuItem
            icon="feed"
            name="feed"
            active={currentUrlPath === '/feed'}
          />
        </div>
      </Link>
      <Link href="/subscriptions" scroll={false}>
        <div>
          <MenuItem
            icon="heart"
            name="subscriptions"
            active={currentUrlPath === '/subscriptions'}
          />
        </div>
      </Link>
    </div>
  )
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    userSignedIn: getIsUserSignedIn(state),
    currentUrlPath: getCurrentUrlPath(state),
  }
}

export default connect<StateToProps, {}, {}, AppState>(mapStateToProps)(
  NavbarBottom,
)
