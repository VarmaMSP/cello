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

const MenuItem: React.SFC<{
  icon: Icon
  name: string
  href: string
  active?: boolean
}> = ({ icon, name, href, active }) => {
  const Icon = iconMap[icon]
  return (
    <div className="w-1/3 flex-none text-center cursor-pointer">
      <Link href={href} scroll={false}>
        <a>
          <Icon
            className={classNames('w-5 h-5 mx-auto fill-current', {
              'text-gray-700': !active,
              'text-green-600': active,
            })}
          />
          <h4 className="capitalize text-sm leading-loose tracking-wide">
            {name}
          </h4>
        </a>
      </Link>
    </div>
  )
}

const NavbarBottom: React.SFC<StateToProps> = ({ currentUrlPath }) => {
  return (
    <div className="flex h-full pt-2 pb-0 bg-white z-50">
      <MenuItem
        icon="home"
        name="home"
        href="/"
        active={currentUrlPath === '/'}
      />
      <MenuItem
        icon="feed"
        name="feed"
        href="/feed"
        active={currentUrlPath === '/feed'}
      />
      <MenuItem
        icon="heart"
        name="subscriptions"
        href="/subscriptions"
        active={currentUrlPath === '/subscriptions'}
      />
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
