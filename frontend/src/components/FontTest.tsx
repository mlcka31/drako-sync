export default function FontTest() {
  return (
    <div className="space-y-4 p-4">
      <h1 className="text-3xl">This should use Superstar font by default</h1>
      <p className="font-geist">This uses Geist font explicitly</p>
      <p className="font-superstarOrig">This uses SuperstarOrig font</p>
    </div>
  );
}
