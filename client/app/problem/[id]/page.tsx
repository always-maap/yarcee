interface Question {
  id: number;
  no: number;
  name: string;
  subject: string;
  difficulty: string;
  problem: string;
  solution: string;
}

async function getQuestion(id: string) {
  const response = await fetch(`http://localhost:8082/api/v1/questions/${id}`);
  return response.json() as Promise<Question>;
}

interface PageParams {
  id: string;
}

export default async function Page({ params }: { params: PageParams }) {
  const question = await getQuestion(params.id);
  return (
    <div>
      <h1>
        {question.no} {question.name}
      </h1>
      <p className="whitespace-pre-wrap">{question.problem}</p>
    </div>
  );
}
