import React from 'react'
import classnames from 'classnames'

const SignInButton: React.SFC<{}> = () => (
  <button className="md:w-24 w-20 h-8 lg:mt-2 bg-orange-600 hover:bg-orange-700 active:bg-orange-700 rounded-lg focus:outline-none focus:shadow-outline ">
    <p className="text-sm text-white font-semibold leading-loose">SIGN IN</p>
  </button>
)

export default SignInButton
