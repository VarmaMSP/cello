import classNames from 'classnames'
import { Icon, iconMap } from 'components/icon'
import { Link } from 'components/link'
import { useRouter } from 'next/router'
import React from 'react'

const MenuItem: React.FC<{
  icon: Icon
  name: string
  href: string
  active?: boolean
}> = ({ icon, name, href, active }) => {
  const Icon = iconMap[icon]
  return (
    <div className="w-1/4 flex-none text-center cursor-pointer">
      <Link href={href} scroll={false}>
        <a>
          <Icon
            className={classNames('w-4 h-4 mx-auto fill-current', {
              'text-gray-700': !active,
              'text-yellow-600': active,
            })}
          />
          <h4 className={classNames('capitalize text-sm leading-loose tracking-wide', {
              'text-gray-700': !active,
              'text-yellow-600 font-medium': active,
            })}
          >
            {name}
          </h4>
        </a>
      </Link>
    </div>
  )
}

const NavbarBottom: React.FC<{}> = () => {
  const currentUrlPath = useRouter().asPath

  return (
    <div className="flex h-full pt-2 pb-0 bg-white z-50">
      <MenuItem
        icon="explore"
        name="explore"
        href="/"
        active={currentUrlPath === '/'}
      />
      <MenuItem
        icon="heart"
        name="subs"
        href="/subscriptions"
        active={currentUrlPath === '/subscriptions'}
      />
      <MenuItem
        icon="history"
        name="history"
        href="/history"
        active={currentUrlPath === '/history'}
      />
      <MenuItem
        icon="playlist"
        name="playlists"
        href="/playlists"
        active={currentUrlPath === '/playlists'}
      />
    </div>
  )
}

export default NavbarBottom
