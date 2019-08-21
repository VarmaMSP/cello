import React from 'react'
import classnames from 'classnames'

const SignInButton: React.SFC<{}> = () => (
  <button className="w-20 h-8 text-blue-600 border-2 border-orange-600 rounded-lg outline-none">
    <p className="w-full text-sm text-orange-600 font-semibold leading-loose text-center">SIGN IN</p>
  </button>
)

export default SignInButton