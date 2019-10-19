import ButtonSignin from 'components/button_signin'
import { iconMap } from 'components/icon'
import Link from 'next/link'
import React from 'react'
import MenuItem from './components/menu_item'

export interface StateToProps {
  userSignedIn: boolean
  currentUrlPath: string
}

const NavbarSide: React.SFC<StateToProps> = (props) => {
  const { userSignedIn, currentUrlPath } = props
  const LogoIcon = iconMap['phenopod']

  return (
    <div className="fixed left-0 top-0 lg:flex flex-col hidden h-screen w-56 px-3 bg-white shadow">
      <LogoIcon className="w-14 h-14 mx-auto mt-2 mb-5" />
      <ul>
        <Link href="/" scroll={false}>
          <li className="h-10 my-1">
            <MenuItem icon="home" name="home" active={currentUrlPath === '/'} />
          </li>
        </Link>
        <Link href="/feed" scroll={false}>
          <li className="h-10 my-1">
            <MenuItem
              icon="feed"
              name="feed"
              active={currentUrlPath === '/feed'}
            />
          </li>
        </Link>
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
            Subscribe to podcasts, create playlists and much more
          </div>
          <div className="w-4/5 h-8 mx-auto">
            <ButtonSignin />
          </div>
        </div>
      ) : (
        <></>
      )}
    </div>
  )
}

export default NavbarSide
