import React from 'react'
import Feed from './feed'
import Subscriptions from './subscriptions'

const SubscriptionsView: React.FC<{}> = () => {
  return (
    <div className="flex md:flex-row flex-col">
      <div className="lg:w-2/3 w-full lg:mr-5">
        <Feed />
      </div>
      <div className="lg:w-1/3 w-full lg:ml-3">
        <Subscriptions />
      </div>
    </div>
  )
}

export default SubscriptionsView
