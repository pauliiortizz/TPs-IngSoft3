export default function Page() {
  return (
    <div className="min-h-screen bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-4xl mx-auto">
        <div className="text-center mb-12">
          <h1 className="text-4xl font-bold text-gray-900 mb-4">WebLearn Docker Project</h1>
          <p className="text-xl text-gray-600">Containerized Learning Platform with QA and Production Environments</p>
        </div>

        <div className="grid md:grid-cols-2 gap-8 mb-12">
          <div className="bg-white rounded-lg shadow-md p-6">
            <h2 className="text-2xl font-semibold text-blue-600 mb-4">QA Environment</h2>
            <div className="space-y-2">
              <p>
                <strong>Frontend:</strong>{" "}
                <a href="http://localhost:8001" className="text-blue-500 hover:underline">
                  localhost:8001
                </a>
              </p>
              <p>
                <strong>Backend API:</strong>{" "}
                <a href="http://localhost:8081/cursos" className="text-blue-500 hover:underline">
                  localhost:8081
                </a>
              </p>
              <p>
                <strong>Database:</strong> MySQL on port 3308
              </p>
              <p>
                <strong>Mode:</strong> Debug enabled
              </p>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6">
            <h2 className="text-2xl font-semibold text-green-600 mb-4">Production Environment</h2>
            <div className="space-y-2">
              <p>
                <strong>Frontend:</strong>{" "}
                <a href="http://localhost:8002" className="text-blue-500 hover:underline">
                  localhost:8002
                </a>
              </p>
              <p>
                <strong>Backend API:</strong>{" "}
                <a href="http://localhost:8082/cursos" className="text-blue-500 hover:underline">
                  localhost:8082
                </a>
              </p>
              <p>
                <strong>Database:</strong> MySQL on port 3309
              </p>
              <p>
                <strong>Mode:</strong> Production optimized
              </p>
            </div>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow-md p-6 mb-8">
          <h2 className="text-2xl font-semibold text-gray-800 mb-4">Docker Architecture</h2>
          <div className="grid md:grid-cols-3 gap-6">
            <div className="text-center">
              <div className="bg-blue-100 rounded-full w-16 h-16 flex items-center justify-center mx-auto mb-3">
                <span className="text-2xl">🐳</span>
              </div>
              <h3 className="font-semibold">Multi-stage Builds</h3>
              <p className="text-sm text-gray-600">Optimized images with minimal size</p>
            </div>
            <div className="text-center">
              <div className="bg-green-100 rounded-full w-16 h-16 flex items-center justify-center mx-auto mb-3">
                <span className="text-2xl">🔄</span>
              </div>
              <h3 className="font-semibold">Environment Separation</h3>
              <p className="text-sm text-gray-600">Isolated QA and PROD environments</p>
            </div>
            <div className="text-center">
              <div className="bg-purple-100 rounded-full w-16 h-16 flex items-center justify-center mx-auto mb-3">
                <span className="text-2xl">📦</span>
              </div>
              <h3 className="font-semibold">Version Control</h3>
              <p className="text-sm text-gray-600">Semantic versioning with v1.0 tags</p>
            </div>
          </div>
        </div>

        <div className="bg-gray-800 text-white rounded-lg p-6">
          <h2 className="text-xl font-semibold mb-4">Quick Start Commands</h2>
          <div className="space-y-2 font-mono text-sm">
            <p>
              <span className="text-green-400">$</span> docker-compose up -d
            </p>
            <p>
              <span className="text-green-400">$</span> docker-compose ps
            </p>
            <p>
              <span className="text-green-400">$</span> docker-compose logs -f
            </p>
          </div>
        </div>
      </div>
    </div>
  )
}
