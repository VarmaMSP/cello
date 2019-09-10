import classNames from 'classnames'
import { iconMap } from 'components/icon'
import { WithRouterProps } from 'next/dist/client/with-router'
import Link from 'next/link'
import { withRouter } from 'next/router'
import React from 'react'

const NavbarBottom: React.SFC<WithRouterProps> = ({ router }) => {
  const IconExplore = iconMap['explore']
  const IconFeed = iconMap['feed']
  const IconSubscriptions = iconMap['heart']
  const IconProfile = iconMap['user-solid-circle']
  let pathname = !!router ? router.pathname : ''

  return (
    <section className="flex justify-around h-full pt-1 pb-0 bg-white z-50">
      <Link href="/">
        <button
          className={classNames('w-20 rounded-full focus:outline-none', {
            'text-green-600': pathname === '',
            'text-gray-700': pathname !== '',
          })}
        >
          <IconExplore className="w-5 h-5 mx-auto fill-current" />
          <h4 className="capitalize text-xs text-center leading-loose">
            discover
          </h4>
        </button>
      </Link>
      <button className="w-20 text-gray-700 rounded-full focus:outline-none">
        <IconFeed className="w-5 h-5 mx-auto fill-current" />
        <h4 className="capitalize text-xs text-center leading-loose">feed</h4>
      </button>
      <button className="w-20 text-gray-700 rounded-full focus:outline-none">
        <IconSubscriptions className="w-5 h-5 mx-auto fill-current" />
        <h4 className="capitalize text-xs text-center leading-loose">
          subscriptions
        </h4>
      </button>
      <button className="w-20 text-gray-700 rounded-full focus:outline-none">
        <IconProfile className="w-5 h-5 mx-auto fill-current" />
        <h4 className="capitalize text-xs text-center leading-loose">
          profile
        </h4>
      </button>
    </section>
  )
}

export default withRouter(NavbarBottom)
