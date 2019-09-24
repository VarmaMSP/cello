import { iconMap } from 'components/icon'
import SignInButton from 'components/sign_in_button'
import Link from 'next/link'
import React from 'react'
import MenuItem from './components/menu_item'

export interface StateToProps {
  userSignedIn: boolean
  currentUrlPath: string
}

const NavbarSide: React.SFC<StateToProps> = (props) => {
  const { userSignedIn, currentUrlPath } = props
  const LogoIcon = iconMap['logo-md']

  return (
    <div className="fixed left-0 top-0 lg:flex flex-col hidden h-screen w-56 px-3 bg-white shadow">
      <LogoIcon className="mx-auto mt-4 mb-8" />
      <ul>
        <Link href="/" scroll={false}>
          <li className="h-10 my-1">
            <MenuItem
              icon="explore"
              name="home"
              active={currentUrlPath === '/'}
            />
          </li>
        </Link>
        <li className="h-10 my-1">
          <MenuItem
            icon="feed"
            name="feed"
            active={currentUrlPath === '/feed'}
          />
        </li>
        <Link href="/subscriptions" scroll={false}>
          <li className="h-10 my-1">
            <MenuItem
              icon="heart"
              name="subscriptions"
              active={currentUrlPath === '/subscriptions'}
            />
          </li>
        </Link>
      </ul>
      <hr className="my-4" />
      {!userSignedIn ? (
        <div className="w-full my-3">
          <div className="mb-2 text-center text-sm text-gray-800 tracking-tighter">
            Subscribe to podcasts, curate episodes and do much more
          </div>
          <div className="w-4/5 h-8 mx-auto">
            <SignInButton />
          </div>
        </div>
      ) : (
        <></>
      )}
    </div>
  )
}

export default NavbarSide
