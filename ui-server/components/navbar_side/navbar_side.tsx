import { iconMap } from 'components/icon'
import { WithRouterProps } from 'next/dist/client/with-router'
import Link from 'next/link'
import { withRouter } from 'next/router'
import React from 'react'
import MenuItem from './components/menu_item'

const NavbarSide: React.SFC<WithRouterProps> = ({ router }) => {
  const pathname = !!router ? router.pathname : ''
  const LogoIcon = iconMap['logo-md']

  return (
    <div className="fixed left-0 top-0 lg:flex flex-col hidden h-screen w-56 px-3 bg-white shadow">
      <LogoIcon className="mx-auto mt-4 mb-8" />
      <ul>
        <Link href="/">
          <li className="my-3">
            <MenuItem icon="explore" name="discover" active={pathname === ''} />
          </li>
        </Link>
        <li className="my-3">
          <MenuItem icon="feed" name="feed" />
        </li>
        <li className="my-3">
          <MenuItem icon="heart" name="subscriptions" />
        </li>
        <li className="my-3">
          <MenuItem icon="user-solid-circle" name="profile" />
        </li>
      </ul>
    </div>
  )
}

export default withRouter(NavbarSide)
