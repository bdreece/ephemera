import Header from "./header";
import Footer from "./footer";
import Sidebar from "./sidebar";

export default function Layout() {
  return (
    <>
      <div>
        <Header />
        <main></main>
        <Footer />
      </div>

      <Sidebar />
    </>
  );
}
