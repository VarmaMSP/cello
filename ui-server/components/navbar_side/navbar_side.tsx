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
    <div className="fixed left-0 top-0 lg:flex hidden flex-col justify-between h-screen w-56 px-3 bg-white border-r">
      <div>
        <LogoIcon className="w-14 h-14 mx-auto mt-2 mb-5" />
        <ul>
          <Link href="/" scroll={false}>
            <li className="h-8">
              <MenuItem
                icon="home"
                name="home"
                active={currentUrlPath === '/'}
              />
            </li>
          </Link>
          <Link href="/feed" scroll={false}>
            <li className="h-8">
              <MenuItem
                icon="feed"
                name="feed"
                active={currentUrlPath === '/feed'}
              />
            </li>
          </Link>
          <Link href="/subscriptions" scroll={false}>
            <li className="h-8">
              <MenuItem
                icon="heart"
                name="subscriptions"
                active={currentUrlPath === '/subscriptions'}
              />
            </li>
          </Link>
          <Link href="/history" scroll={false}>
            <li className="h-8">
              <MenuItem
                icon="history"
                name="history"
                active={currentUrlPath === '/history'}
              />
            </li>
          </Link>
        </ul>

        {!userSignedIn && (
          <>
            <hr className="my-4" />
            <div className="w-full my-3">
              <div className="mb-2 text-center text-sm text-gray-800 tracking-tighter leading-tight">
                Subscribe to podcasts, create playlists and much more
              </div>
              <div className="w-4/5 h-8 mx-auto">
                <ButtonSignin />
              </div>
            </div>
          </>
        )}
      </div>
      <div className="px-2 py-6 text-sm text-gray-800">
        <p className="leading-tight">
          <Link href="/about" prefetch={false}>
            <a>{'about'}</a>
          </Link>{' '}
          <span className="font-extrabold">&middot;</span>{' '}
          <Link href="/privacy" prefetch={false}>
            <a>{'privacy'}</a>
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
