import NextAuth, { AuthOptions } from "next-auth"
import GithubProvider from "next-auth/providers/github"
import CredentialsProvider from "next-auth/providers/credentials"
import axios from 'axios'

// import bcrypt from 'bcrypt'



export const authOptions: any = {
  // Configure one or more authentication providers
  providers: [
    CredentialsProvider({
      // The name to display on the sign in form (e.g. 'Sign in with...')
      name: 'credentials',
      // The credentials is used to generate a suitable form on the sign in page.
      // You can specify whatever fields you are expecting to be submitted.
      // e.g. domain, username, password, 2FA token, etc.
      // You can pass any HTML attribute to the <input> tag through the object.
      credentials: {
        username: { label: "Username", type: "text" },
        password: { label: "Password", type: "password" }
      },

      async authorize(credentials) {
        // You need to provide your own logic here that takes the credentials
        // submitted and returns either a object representing a user or value
        // that is false/null if the credentials are invalid.
        // e.g. return { id: 1, name: 'J Smith', email: 'jsmith@example.com' }
        // You can also use the `req` object to obtain additional parameters
        // (i.e., the request IP address)
        // console.log("credentials is ===>"  , credentials)

        let data = {
          username: credentials?.username,
          password: credentials?.password
        }

        // console.log("data is ================================>" , data)

        // const res = await axios.post("http://127.0.0.1:8000/signin", data)
        // .then((res) => {
        //   console.log("res from go is ===>" , res.data)
        // }).catch((error) => {
        //   console.log("error is =====>" , error)
        // })

        const res = await fetch("http://127.0.0.1:8000/signin", {
          method: 'POST',
          body: JSON.stringify(data),
          headers: { "Content-Type": "application/json" }
        })
        const user = await res.json()

        if (res.status && user) {
          return user
        } else {
          return null
        }

        // If no error and we have user data, return it
        // if (res.ok && user) {
        //   return user
        // }
        // Return null if user data could not be retrieved
      }
    })],
  secret: "LlKq6ZtYbr+hTC073mAmAh9/h2HwMfsFo4hrfCx5mLg=",
  session: {
    strategy: "jwt"
  },
  // callbacks: {

  //   async jwt({ token, user ,  session }) {

  //     return { ...token, ...user };
  //   },

  //   session({ session, token, user }) {
  //     return session // The return type will match the one returned in `useSession()`
  //   },
  // },



}
export default NextAuth(authOptions)

