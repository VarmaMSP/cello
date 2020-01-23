import { iconMap } from 'components/icon'
import { Link } from 'components/link'
import SearchBar from 'components/search_bar/side_navbar'
import { useRouter } from 'next/router'
import React from 'react'
import { ViewportSize } from 'types/app'
import MenuItem from './components/menu_item'

export interface StateToProps {
  userSignedIn: boolean
  viewportSize: ViewportSize
}

export interface DispatchToProps {
  showSigninModal: () => void
}

const NavbarSide: React.SFC<StateToProps & DispatchToProps> = ({
  userSignedIn,
  viewportSize,
  showSigninModal,
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
        <ul>
          <Link href="/" scroll={false}>
            <li className="h-10">
              <MenuItem
                icon="explore"
                name="explore"
                active={currentUrlPath === '/'}
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

        {!userSignedIn && (
          <>
            <hr className="my-4" />
            <div className="w-full my-3">
              <p className="mb-2 text-center text-xs text-gray-700 tracking-wide leading-normal">
                Discover podcasts, create playlists and much more.{' '}
                <span
                  className="text-red-800 font-medium underline cursor-pointer"
                  onClick={() => showSigninModal()}
                >
                  sign in
                </span>
              </p>
            </div>
          </>
        )}
      </div>
      <div className="px-2 py-6 text-sm text-gray-800">
        <p className="leading-tight">
          <Link href="/about" prefetch={false}>
            <a className="cursor-pointer">{'about'}</a>
          </Link>{' '}
          <span className="font-extrabold">&middot;</span>{' '}
          <Link href="/privacy" prefetch={false}>
            <a className="cursor-pointer">{'privacy'}</a>
          </Link>
        </p>
        <p className="mb-1">
          <a href="https://www.facebook.com/phenopod" target="_blank">
            {'facebook'}
          </a>{' '}
          <span className="font-extrabold">&middot;</span>{' '}
          <a href="https://twitter.com/phenopod" target="_blank">
            {'twitter'}
          </a>{' '}
          <span className="font-extrabold">&middot;</span>{' '}
          <a href="https://www.reddit.com/r/phenopod/" target="_blank">
            {'reddit'}
          </a>
        </p>
        <a href="mailto:hello@phenopod.com" className="font-light">
          {'hello@phenopod.com'}
        </a>
      </div>
    </div>
  )
}

export default NavbarSide
