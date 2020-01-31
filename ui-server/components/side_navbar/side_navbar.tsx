import About from 'components/about'
import { iconMap } from 'components/icon'
import SearchBar from 'components/search_bar/side_navbar'
import SignInButton from 'components/sign_in_button'
import { useRouter } from 'next/router'
import React from 'react'
import { ViewportSize } from 'types/app'
import MenuItem from './components/menu_item'

export interface StateToProps {
  userSignedIn: boolean
  viewportSize: ViewportSize
}

const NavbarSide: React.SFC<StateToProps> = ({
  userSignedIn,
  viewportSize,
}) => {
  if (viewportSize === 'SM') {
    return <></>
  }

  const LogoIcon = iconMap['phenopod']
  const currentUrlPath = useRouter().asPath

  return (
    <div className="fixed left-0 top-0 flex flex-col justify-between h-screen w-64 pl-4 pr-2 bg-white">
      <div>
        <LogoIcon className="w-14 h-14 mx-auto mt-2 mb-3" />
        <div className="mb-6">
          <SearchBar />
        </div>
        <ul className="mb-10">
          <li className="h-10">
            <MenuItem icon="home" name="home" href="/" />
          </li>
          <li className="h-10">
            <MenuItem icon="explore" name="explore" href="/explore" />
          </li>
          <li className="h-10">
            <MenuItem icon="heart" name="subscriptions" href="/subscriptions" />
          </li>
          <li className="h-10">
            <MenuItem icon="history" name="history" href="/history" />
          </li>
          <li className="h-10">
            <MenuItem icon="playlist" name="playlists" href="/playlists" />
          </li>
        </ul>

        {currentUrlPath !== '/' && !userSignedIn && (
          <div className="h-10 px-5">
            <SignInButton />
          </div>
        )}
      </div>
      <div className="px-2 py-6">
        <About />
      </div>
    </div>
  )
}

export default NavbarSide
