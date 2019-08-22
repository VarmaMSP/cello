import React from 'react'
import classnames from 'classnames'

const SignInButton: React.SFC<{}> = () => (
  <button className="md:w-24 w-20 h-8 lg:mt-2 bg-orange-600 text-white rounded-lg outline-none">
    <p className="w-full text-sm font-semibold leading-loose text-center">SIGN IN</p>
  </button>
)

export default SignInButton