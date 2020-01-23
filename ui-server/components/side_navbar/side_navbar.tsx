import About from 'components/about'
import { iconMap } from 'components/icon'
import { Link } from 'components/link'
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
          <Link href="/" scroll={false}>
            <li className="h-10">
              <MenuItem
                icon="home"
                name="home"
                active={currentUrlPath === '/'}
              />
            </li>
          </Link>

          <Link href="/explore" scroll={false}>
            <li className="h-10">
              <MenuItem
                icon="explore"
                name="explore"
                active={currentUrlPath === '/explore'}
              />
            </li>
          </Link>

          <Link href="/subscriptions" scroll={false}>
            <li className="h-10">
              <MenuItem
                icon="heart"
                name="subscriptions"
                active={currentUrlPath === '/subscriptions'}
              />
            </li>
          </Link>

          <Link href="/history" scroll={false}>
            <li className="h-10">
              <MenuItem
                icon="history"
                name="history"
                active={currentUrlPath === '/history'}
              />
            </li>
          </Link>

          <Link href="/playlists" scroll={false}>
            <li className="h-10">
              <MenuItem
                icon="playlist"
                name="playlists"
                active={currentUrlPath === '/playlists'}
              />
            </li>
          </Link>
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
