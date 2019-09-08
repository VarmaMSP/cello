import { iconMap } from 'components/icon'
import React from 'react'

const NavbarBottom: React.SFC<{}> = () => {
  const IconExplore = iconMap['explore']
  const IconFeed = iconMap['feed']
  const IconSubscriptions = iconMap['heart']
  const IconProfile = iconMap['user-solid-circle']

  return (
    <section className="flex justify-around h-full pt-1 pb-0 bg-white z-50">
      <button className="w-20 text-green-600 rounded-full focus:outline-none">
        <IconExplore className="w-5 h-5 mx-auto fill-current" />
        <h4 className="capitalize text-xs text-center leading-loose">
          discover
        </h4>
      </button>
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

export default NavbarBottom
